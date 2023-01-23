import { getBaseUrl } from "$lib/utils/urls";

export const load = async () => {
  const res = await fetch(getBaseUrl() + `/api/cocktails-thebar.json`);
  let cocktails = await res.json();

  return { cocktails };
}