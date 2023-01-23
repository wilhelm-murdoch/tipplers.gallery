<script lang="ts">
	import type { NavigationItem } from './utils';

	export let items: NavigationItem[] = [];
	export let active: string;
</script>

<div class="max-w-7xl mx-auto">
	<nav class="-mb-px flex justify-between" aria-label="Tabs">
		{#each items as item (item.slug)}
			{#if item.slug === active}
				<a href={item.path} title={item.title} class="border-tipplers-primary text-tipplers-primary py-4 px-3 text-center border-b-4 font-bold text-sm md:text-md" aria-current="page">{item.name}</a>
			{:else}
				<a href={item.path} title={item.title} class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 py-4 px-3 text-center border-b-2 font-medium text-sm md:text-md">{item.name}</a>
			{/if}
		{/each}
	</nav>
</div>
