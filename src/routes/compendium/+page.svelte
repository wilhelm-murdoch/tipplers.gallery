<script lang="ts">
	import { CardSlim, Footer, Header, NavigationFilter, NavigationSub } from '$components';
	import type { NavigationItem } from '$components/Navigation/utils';
	import type { Cocktails, Cocktail } from '$lib/utils/types';
	import { onMount } from 'svelte';
	import SvelteSeo from 'svelte-seo';
	import { fade } from 'svelte/transition';
	import { navigationItems } from './utils';

	export let data: Cocktails;

	let alphaNumericFilterItems: NavigationItem[] = [];
	let filtered: Cocktail[] = [];
	let activeFilter: string = 'a';

	const handleFilterItemClick = (e: any) => {
		activeFilter = e.detail.slug;
		filtered = data.cocktails.filter((v) => v.name[0].toLowerCase() === e.detail.slug);
	};

	onMount(() => {
		filtered = data.cocktails.filter((v) => v.name[0].toLowerCase() === activeFilter);

		alphaNumericFilterItems = '#abcdefghijklmnopqrstuvwxyz'.split('').map((c) => {
			return {
				name: c.toUpperCase(),
				slug: c,
				path: '#',
				title: 'Filter cocktails starting with "' + c + '".',
				count: data.cocktails.filter((v) => v.name[0].toLowerCase() === c).length
			};
		});
	});

	const twitterMeta = {
		site: '@wilhelm',
		title: "Tippler's Gallery ~ Compendium",
		description: 'Quickly search & browse through thousands of cocktail recipes gathered from around the web.',
		image: 'https://tipplers.gallery/images/og-cover.png',
		imageAlt: "Tippler's Gallery"
	};

	const openGraphMeta = {
		title: "Tippler's Gallery ~ Compendium",
		description: 'Quickly search & browse through thousands of cocktail recipes gathered from around the web.',
		type: 'article',
		url: 'https://tipplers.gallery/compendium',
		images: [
			{
				url: 'https://tipplers.gallery/images/og-cover.png',
				alt: "Tippler's Gallery"
			}
		]
	};
</script>

<div class="px-2">
	<Header />
</div>

<div class="max-w-7xl mx-auto text-center py-10">
	<h2 class="text-6xl font-thin italic text-tipplers-secondary">The Compendium</h2>
</div>

<NavigationSub items={navigationItems} active="a-z" />
<NavigationFilter items={alphaNumericFilterItems} active={activeFilter} on:click={handleFilterItemClick} />

<div class="max-w-7xl mx-auto px-2">
	{#if filtered.length}
		<section>
			<h3 id={activeFilter} class="py-8 text-2xl font-semibold font-serif border-b border-slate-50 bg-tipplers-lightest/25">
				<span class="underline">{activeFilter.toUpperCase()}</span>
				<span class="bg-white px-2 py-1 rounded shadow-sm text-xs font-sans align-top">
					{filtered.length}
				</span>
			</h3>
			<div class="grid grid-cols-2 gap-6 md:grid-cols-4 lg:grid-cols-6">
				{#each filtered as cocktail, i (cocktail.cid)}
					<div in:fade={{ delay: i * 100 }}>
						<CardSlim name={cocktail.name} src={cocktail.images[0].relative_path} from={cocktail.meta.sourced_from} attribution={cocktail.images[0].attribution} url={cocktail.meta.source_url} />
					</div>
				{/each}
			</div>
		</section>
	{/if}
</div>

<Footer />

<SvelteSeo title="Tippler's Gallery ~ Compendium" description="Quickly search & browse through thousands of cocktail recipes gathered from around the web." canonical="https://tipplers.gallery/compendium" twitter={twitterMeta} openGraph={openGraphMeta} />
