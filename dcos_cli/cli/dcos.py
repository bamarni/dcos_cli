"""The CLI for DC/OS."""

import click

from dcos_cli import __version__
from dcos_cli.cli.cluster import cluster


def print_version(ctx, param, value):
    """Print the DC/OS CLI version."""
    # pylint: disable=unused-argument
    if not value or ctx.resilient_parsing:
        return
    click.echo(__version__)
    ctx.exit()


class DCOSCLI(click.MultiCommand):
    """The main `dcos` command."""

    def list_commands(self, ctx):
        """Return the list of available DC/OS commands."""
        return ['cluster']

    def get_command(self, ctx, cmd_name):
        """Get a given command."""
        return cluster


@click.command(cls=DCOSCLI)
@click.option('--debug', is_flag=True, help="Enable debug mode.")
@click.option(
    '--version',
    is_flag=True,
    callback=print_version,
    expose_value=False,
    is_eager=True,
    help="Print version information.",
)
def dcos(debug):
    """Run the dcos command."""
    assert isinstance(debug, bool)
