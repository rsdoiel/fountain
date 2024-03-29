#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = fdx

PANDOC = $(shell which pandoc)

MD_PAGES = $(shell ls -1 *.md | grep -v 'nav.md')

HTML_PAGES = $(shell ls -1 *.md | grep -v 'nav.md' | sed -E 's/.md/.html/g')

build: $(HTML_PAGES) $(MD_PAGES) pagefind

$(HTML_PAGES): $(MD_PAGES) .FORCE
	if [ -f $(PANDOC) ]; then $(PANDOC) --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
		--lua-filter=links-to-html.lua \
	    --template=page.tmpl; git add $(basename $@).html; fi
	@if [ $@ = "README.html" ]; then git mv README.html index.html; fi

pagefind: .FORCE
	#pagefind --verbose --exclude-selectors="nav,header,footer" --bundle-dir ./pagefind --source .
	pagefind --verbose --exclude-selectors="nav,header,footer" --site .
	git add pagefind

clean:
	@if [ -f index.html ]; then rm *.html; fi

.FORCE:
