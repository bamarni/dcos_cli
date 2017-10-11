# DC/OS CLI

The DC/OS Command Line Interface (CLI) is a cross-platform command line utility that provides a user-friendly yet powerful way to manage DC/OS clusters.

## Installation

## Usage

## Development environment

This package requires Python 3.5 and pip. We recommend setting-up a virtual environment, there are various ways to do it.
The simplest one being to use the built-in [venv](https://docs.python.org/3/library/venv.html) module :

    # Creates a virtualenv inside an env directory
    python3.5 -m venv env
    # Activates the virtualenv
    source bin/env/activate

Then you need to install development dependencies :

    pip install -e .[dev]

You should now be able to invoke the `dcos` command.

## Running tests
