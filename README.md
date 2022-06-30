# Gintokipona CLI

My buddy and I made a deal that I would watch [Gintama](https://jisho.org/search/kintama)
if he learns [Toki Pona](https://en.wikipedia.org/wiki/Toki_Pona).

Being a programmer, I just had to overengineer a system for keeping track
of how much we've each contributed to the deal.

## Usage
Run the executable in bin with one of the following commands:

 - `list` - Shows which Actions are available
 - `stats <action>` - Shows the stats for a given action
 - `record <action>` - Records that you fulfilled a certain action

Actions are things to be fulfilled over time, e.g. "Watch Gintama"

The records are stored in a JSON file either named `ginponadata.json`
or what ever you specify with the `-f` option. An example is included

That's about it

If you want more features, they should be relatively easy to implement
so just open an issue
