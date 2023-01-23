import { getBaseUrl } from "$lib/utils/urls";

export const load = async () => {
  const response = await fetch(getBaseUrl() + `/api/tags.json`);

  let tags = await response.json();

  if (response.status === 200) {
    return { tags };
  }

  throw new Error(response.statusText);
}