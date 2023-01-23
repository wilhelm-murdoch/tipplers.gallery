export interface NavigationItem {
  name: string;
  slug: string;
  path: string;
  title: string;
  count?: number;
}

export const emptyNavigationItem = {
  name: '',
  slug: '',
  path: '',
  title: '',
  count: 0
}