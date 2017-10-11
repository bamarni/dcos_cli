import click

from dcos_cli import __version__

def print_version(ctx, param, value):
    if not value or ctx.resilient_parsing:
        return
    click.echo(__version__)
    ctx.exit()

@click.group()
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
    pass

import dcos_cli.cli.cluster # noqa: E402
