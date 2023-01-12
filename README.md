Make Me A Cocktial

```
curl 'https://makemeacocktail.com/scripts/ajaxcalls.php?ajaxcall=updateResults' \
  --data-raw 'ajaxurl=scripts/ajaxcalls.php?ajaxcall=updateResults&scid=885300&numPerPage=10&sortBy=c_score&pageNum=2000' \
  --compressed
```
Grab the cocktail id from the `data-cid=` attribute. Use the value of this attribute in the following URL structure:
```
https://makemeacocktail.com/cocktail/{id}/cocktail-name/
```
`cocktail-name` can literally be anything. The id is the canonical path to the target page; follow the redirect.

Difford's Guide
https://www.diffordsguide.com/en-au/cocktails/search?ingredient%5B%5D=g-58&ingredient%5B%5D=g-68&ingredient%5B%5D=g-1&ingredient%5B%5D=g-52&ingredient%5B%5D=g-63&ingredient%5B%5D=g-23&ingredient%5B%5D=g-124&ingredient%5B%5D=g-142&include%5Bdg%5D=1&limit=80&sort=rating&offset=80

