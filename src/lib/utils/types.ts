export type Sluggable = {
  name: string;
  slug: string;
};

export type Image = {
  source_url: string;
  attribution: string;
  relative_path: string;
};

export type Ingredient = {
  name: string;
  measurement: number;
  unit: string;
}

export type Meta = {
  standard_dinks: number;
  sourced_from: string;
  source_url: string;
}

export interface Cocktail {
  cid: string;
  name: string;
  description: string;
  ingredients: Ingredient[];
  instructions: string[];
  images: Image[];
  equipment: Sluggable[];
  tags: Sluggable[];
  meta: Meta;
}

export interface Cocktails {
  cocktails: Cocktail[]
}