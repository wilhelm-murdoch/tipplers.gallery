<script lang="ts">
	import { NavigationSub, NavigationFilter } from '$components';
	import type { NavigationItem } from '$components/Navigation/utils';
	import type { Tag, Tags } from '$lib/utils/types';
	import { onMount } from 'svelte';
	import { navigationItems } from '../utils';

	export let data: Tags;

	interface GroupedTags {
		name: string;
		slug: string;
		tags: Tag[];
	}

	let groupedTags: GroupedTags[] = [];
	let alphaNumericFilterItems: NavigationItem[] = [];
	let activeFilter: string;
	let alphaNumericSplit: string[] = '#abcdefghijklmnopqrstuvwxyz'.split('');

	onMount(() => {
		groupedTags = alphaNumericSplit.map((c) => {
			return {
				name: c.toUpperCase(),
				slug: c,
				tags: data.tags.filter((v) => v.name[0].toLowerCase() === c)
			};
		});

		alphaNumericFilterItems = alphaNumericSplit.map((c) => {
			var group = groupedTags.filter((f) => f.slug === c);
			var count = 0;

			if (group) {
				count = group[0].tags.length;
			}

			if (!activeFilter && count) {
				activeFilter = c;
			}

			return {
				name: c.toUpperCase(),
				slug: c,
				path: '#' + c,
				title: 'Jump to ' + c,
				count: count
			};
		});
	});
</script>

<NavigationSub items={navigationItems} active="tags" />
<NavigationFilter items={alphaNumericFilterItems} active={activeFilter} />

<div class="max-w-7xl mx-auto px-2">
	{#each groupedTags as group (group.name)}
		{#if group.tags.length}
			<section>
				<h3 id={group.slug} class="py-8 text-2xl font-semibold font-serif border-b border-slate-50 bg-tipplers-lightest/25">
					"{group.name}"
					<span class="items-center justify-center align-top text-sm font-semibold rounded-md shadow-sm bg-white px-2 py-0.5 font-sans">{group.tags.length}</span>
					<a href="#top" title="Scroll back to the top of the page." class="inline-block p-0.5 align-top text-sm cursor-pointer hover:text-tipplers-secondary">
						<svg class="h-4 w-4" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
							<path stroke-linecap="round" stroke-linejoin="round" d="M12 19.5v-15m0 0l-6.75 6.75M12 4.5l6.75 6.75" />
						</svg>
					</a>
				</h3>
				<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
					{#each group.tags as tag, i (tag)}
						<div class="prose py-1">
							<a href="/compendium/tags/{tag.slug}" title="List cocktails using the {tag.name} tag." class="truncate">{tag.name}</a>
							<span class="items-center justify-center text-xs font-semibold rounded-md shadow-sm bg-white px-2 py-0.5">{tag.count}</span>
						</div>
					{/each}
				</div>
			</section>
		{/if}
	{/each}
</div>
