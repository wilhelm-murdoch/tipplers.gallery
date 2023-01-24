<script lang="ts">
	import { NavigationSub, NavigationFilter } from '$components';
	import type { NavigationItem } from '$components/Navigation/utils';
	import type { Tag, Tags } from '$lib/utils/types';
	import { onMount } from 'svelte';
	import { navigationItems } from '../utils';

	export let data: Tags;

	let alphaNumericFilterItems: NavigationItem[] = [];
	let filtered: Tag[] = [];
	let activeFilter: string = 'a';

	const handleFilterItemClick = (e: any) => {
		activeFilter = e.detail.slug;
		filtered = data.tags.filter((v) => v.name[0].toLowerCase() === e.detail.slug);
	};

	onMount(() => {
		filtered = data.tags.filter((v) => v.name[0].toLowerCase() === activeFilter);

		alphaNumericFilterItems = '#abcdefghijklmnopqrstuvwxyz'.split('').map((c) => {
			return {
				name: c.toUpperCase(),
				slug: c,
				path: '',
				title: 'Filter cocktails starting with "' + c + '".',
				count: data.tags.filter((v) => v.name[0].toLowerCase() === c).length
			};
		});
	});
</script>

<NavigationSub items={navigationItems} active="tags" />
<NavigationFilter items={alphaNumericFilterItems} active={activeFilter} on:click={handleFilterItemClick} />

<div class="max-w-7xl mx-auto px-2">
	{#if filtered.length}
		<section>
			<h3 id={activeFilter} class="py-8 text-2xl font-semibold font-serif border-b border-slate-50 bg-tipplers-lightest/25">
				{filtered.length} found
			</h3>
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
				{#each filtered as tag, i (tag)}
					<div class="prose py-1">
						<a href="/compendium/tags/{tag.slug}" title="List cocktails using the {tag.name} tag." class="truncate">{tag.name}</a>
						<span class="items-center justify-center text-xs font-semibold rounded-md shadow-sm bg-white px-2 py-0.5">{tag.count}</span>
					</div>
				{/each}
			</div>
		</section>
	{/if}
</div>
