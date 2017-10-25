"""Defines the `dcos cluster setup` subcommand."""

import click


@click.command()
def setup():
    """Set up a cluster."""
    print('Setup a cluster.')
