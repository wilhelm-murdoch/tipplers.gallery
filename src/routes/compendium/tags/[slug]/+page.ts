import type { Cocktail, Tag } from "$lib/utils/types";
import { getBaseUrl } from "$lib/utils/urls";

export async function load({ fetch, params }) {
  const [cocktailsRequest, tagsRequest] = await Promise.all([
    fetch(getBaseUrl() + `/api/cocktails-thebar.json`),
    fetch(getBaseUrl() + `/api/tags.json`)
  ])
  if (cocktailsRequest.ok && tagsRequest.ok) {
    const cocktails: Cocktail[] = await cocktailsRequest.json()
    const tags: Tag[] = await tagsRequest.json()

    return {
      slug: params.slug,
      tags: tags,
      cocktails: cocktails.filter((c: Cocktail) => c.tags.find((t) => t.slug == params.slug))
    };
  }

  throw new Error([cocktailsRequest.statusText, tagsRequest.statusText].join('\n'));
}