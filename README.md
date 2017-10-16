# DC/OS CLI [![Build Status](https://travis-ci.org/bamarni/dcos_cli.svg?branch=master)](https://travis-ci.org/bamarni/dcos_cli)

The DC/OS Command Line Interface (CLI) is a cross-platform command line utility that provides a user-friendly yet powerful way to manage DC/OS clusters.

## Installation

## Usage

## Development environment

This package requires Python 3.5 and pip. We recommend setting-up a virtual environment, there are various ways to do it.
The simplest one being to use the built-in [venv](https://docs.python.org/3/library/venv.html) module :

    # Creates a virtualenv inside an env directory
    python3.5 -m venv env
    # Activates the virtualenv
    source env/bin/activate

Then you need to install development dependencies :

    make deps

You should now be able to invoke the `dcos` command.

## Running tests

Once you have the development environment setup, you can run tests with the following command :

    make test
