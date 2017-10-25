"""Defines the `dcos cluster remove` subcommand."""

import click


@click.command()
def remove():
    """Remove cluster(s)."""
    print('Remove cluster(s).')
