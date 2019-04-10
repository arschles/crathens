package queue

func ghFetcher(
	ctx context.Context,
	ghCl *github.Client,
	modCh <-chan string,
	nextCh chan<- github.RepositoryTag,
	ticker *time.Ticker,
) {
	for range <-ticker.C {
		switch {
		case <-ctx.Done():
			return
		case mod := <-modCh:
			tags, err := gh.FetchTags(ctx, ghCl, mod)
			if err != nil {
				log.Printf("Error fetching GH tags for %s (%s)", mod, err)
			}
			nextCh <- tags
		}
	}
}
