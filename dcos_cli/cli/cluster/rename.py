"""Defines the `dcos cluster rename` subcommand."""

import click


@click.command()
def rename():
    """Rename a cluster."""
    print('Rename a cluster.')
