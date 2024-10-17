package main

import (
	"context"
	"encoding/csv"
	"os"

	"github.com/ZacxDev/scrapemate"

	"github.com/ZacxDev/scrapemate/adapters/writers/csvwriter"
	"github.com/ZacxDev/scrapemate/quotestoscrapelogin/quotes"
	"github.com/ZacxDev/scrapemate/scrapemateapp"
)

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
		return
	}
	os.Exit(0)
}

func run() error {
	csvWriter := csvwriter.NewCsvWriter(csv.NewWriter(os.Stdout))

	writers := []scrapemate.ResultWriter{
		csvWriter,
	}

	cfg, err := scrapemateapp.NewConfig(writers,
		scrapemateapp.WithInitJob(quotes.NewLoginCRSFToken()),
	)
	if err != nil {
		return err
	}

	app, err := scrapemateapp.NewScrapeMateApp(cfg)
	if err != nil {
		return err
	}

	seedJobs := []scrapemate.IJob{
		quotes.NewQuoteCollectJob("https://quotes.toscrape.com/"),
	}
	return app.Start(context.Background(), seedJobs...)
}
