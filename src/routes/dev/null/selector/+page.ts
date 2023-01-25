import { getBaseUrl } from "$lib/utils/urls";

export const load = async ({ fetch }) => {
  const response = await fetch(getBaseUrl() + `/api/cocktails-thebar.json`);

  let cocktails = await response.json();

  if (response.status === 200) {
    return { cocktails };
  }

  throw new Error(response.statusText);
}