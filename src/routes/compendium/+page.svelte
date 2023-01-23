<script lang="ts">
	import { CardSlim, NavigationFilter, NavigationSub } from '$components';
	import type { NavigationItem } from '$components/Navigation/utils';
	import type { Cocktails, Cocktail } from '$lib/utils/types';
	import { onMount } from 'svelte';
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
				path: '',
				title: 'Filter cocktails starting with "' + c + '".',
				count: data.cocktails.filter((v) => v.name[0].toLowerCase() === c).length
			};
		});
	});
</script>

<NavigationSub items={navigationItems} active="a-z" />
<NavigationFilter items={alphaNumericFilterItems} active={activeFilter} on:click={handleFilterItemClick} />

<div class="max-w-7xl mx-auto px-2">
	{#if filtered.length}
		<section>
			<h3 id={activeFilter} class="py-8 text-2xl font-semibold font-serif border-b border-slate-50 bg-tipplers-lightest/25">
				{filtered.length} found
			</h3>
			<div class="grid grid-cols-2 gap-6 md:grid-cols-4 lg:grid-cols-6">
				{#each filtered as cocktail, i (cocktail.cid)}
					<div in:fade={{ delay: i * 50 }}>
						<CardSlim name={cocktail.name} src={cocktail.images[0].relative_path} from={cocktail.meta.sourced_from} attribution={cocktail.images[0].attribution} url={cocktail.meta.source_url} />
					</div>
				{/each}
			</div>
		</section>
	{/if}
</div>
