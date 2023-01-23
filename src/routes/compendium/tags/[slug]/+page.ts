import type { Cocktail } from "$lib/utils/types";
import { getBaseUrl } from "$lib/utils/urls";

export const load = async ({ params }) => {
  const response = await fetch(getBaseUrl() + `/api/cocktails-thebar.json`);

  let cocktails = await response.json();

  if (response.status === 200) {
    return {
      cocktails: cocktails.filter((c: Cocktail) => c.tags.find((t) => t.slug == params.slug))
    };
  }

  throw new Error(response.statusText);
}