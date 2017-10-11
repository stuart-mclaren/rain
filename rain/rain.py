#!/usr/bin/env python
# Copyright 2017 Stuart McLaren
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


from datetime import datetime, timedelta, tzinfo
import re
import warnings

import argparse
from dateutil import tz
from lxml import etree
import requests

parser = argparse.ArgumentParser(description='Rain forecast')
parser.add_argument('--hours', metavar='hours', type=int, nargs='?',
                    const=1, default=12, help='Number of hours to forecast.')
parser.add_argument('--longitude', metavar='longitude', type=float,
                    default=-8.9297, help="Longitude. Default: -8.9297")
parser.add_argument('--latitude', metavar='latitude', type=float,
                    default=53.2648, help="Latitude. Default: 53.2648")

args = parser.parse_args()

warnings.filterwarnings("ignore")

URL = "http://api.met.no/weatherapi/locationforecast" \
    "/1.9/?lat=%f;lon=%f" % (args.latitude, args.longitude)


class Zone(tzinfo):
    def __init__(self, offset, isdst, name):
        self.offset = offset
        self.isdst = isdst
        self.name = name

    def utcoffset(self, dt):
        return timedelta(hours=self.offset) + self.dst(dt)

    def dst(self, dt):
        return timedelta(hours=1) if self.isdst else timedelta(0)

    def tzname(self, dt):
        return self.name


def do_forecast():
    GMT = Zone(0, False, 'GMT')
    local_timezone = tz.gettz()
    response = requests.get(URL)
    forecast = etree.fromstring(response.text)

    i = 0
    for _time in forecast.findall(".//time"):
        location = _time.find("location")
        precipitation = location.find("precipitation")
        if precipitation is None:
            continue

        _from = datetime.strptime(_time.get('from'), "%Y-%m-%dT%H:%M:%SZ")
        _from = _from.replace(tzinfo=GMT)
        _to = datetime.strptime(_time.get('to'), "%Y-%m-%dT%H:%M:%SZ")
        _to = _to.replace(tzinfo=GMT)
        num_stars = 10.0 * float(precipitation.get('value'))
        stars = int(num_stars) * '*'
        if _from.hour - _to.hour == 1 or _from.hour - _to.hour == -1:
            i = i + 1
            print("%02d:00 - %02d:00 %s %s | %s" % (
                _from.astimezone(local_timezone).hour,
                _to.astimezone(local_timezone).hour,
                precipitation.get('value'),
                precipitation.get('unit'),
                stars))
        if i == args.hours:
            break


if __name__ == "__main__":
    do_forecast()
