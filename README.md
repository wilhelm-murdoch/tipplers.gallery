# Tippler's Gallery
A compendium of the finest cocktail recipes.

```
mage images:download static/cocktails.json static/images
for image in ./**/original.jpg; do convert "$image" -resize 50% "$(dirname $image)/large.jpg"; done
for image in ./**/original.jpg; do convert "$image" -resize 35% "$(dirname $image)/medium.jpg"; done
for image in ./**/original.jpg; do convert "$image" -resize 250 "$(dirname $image)/thumbnail.jpg"; done
```