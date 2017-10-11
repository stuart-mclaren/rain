from setuptools import setup

setup(
    name='rain',
    packages=['rain'],
    version='0.0.1',
    long_description=__doc__,
    entry_points={
        'console_scripts': ['rain=rain.rain:do_forecast']
    },
    install_requires=[
        "lxml",
        "requests",
        "python-dateutil"
    ],
    include_package_data=True
)
