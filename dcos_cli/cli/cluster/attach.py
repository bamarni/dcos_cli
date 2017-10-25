"""Defines the `dcos cluster attach` subcommand."""

import click


@click.command()
def attach():
    """Attach to a cluster."""
    print('Attach to a cluster.')
