import type { NavigationItem } from "$components/Navigation/utils";

const spirits: string[] = ["Gin", "Whiskey", "Rum", "Cognac", "Absinthe", "Tequila", "Mezcal", "Vodka"]

export const navigationItemsSpirits: NavigationItem[] = spirits.map(v => {
  return {
    name: v,
    slug: v.toLowerCase(),
    title: "An alphabetical listing of all curated " + v + "-based cocktail recipes.",
    path: "/compendium/spirits",
  }
})
