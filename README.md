# Opening hand simulator

This project aims to provide a painless way to simulate opening hands/prizes for the Pokemon TCG.

## Initialising set mappings for decklist parser
Running the command `go run ./api/cmd/initsetmap/main.go` will download and store mappings from ptcgo(Note: NOT ptcg live) set codes to set ids for the decklist parser.
Differences between PTCGO codes and PTCG Live codes have to be manually edited for the time being.
