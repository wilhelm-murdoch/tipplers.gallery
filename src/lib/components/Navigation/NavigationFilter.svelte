<script lang="ts">
	import type { NavigationItem } from './utils';
	import { createEventDispatcher } from 'svelte';

	const wrapperBaseClassList = 'bg-white border-y';

	export let items: NavigationItem[];
	export let active: string;
	export let wrapperClasses: string = '';

	const dispatch = createEventDispatcher();

	const onClickFilterItem = (item: NavigationItem) => {
		active = item.slug;
		dispatch('click', item);
	};

	$: wrapperCompositeClasses = `${wrapperBaseClassList} ${wrapperClasses}`;
</script>

<div class={wrapperCompositeClasses}>
	<div class="max-w-7xl mx-auto">
		<nav class="flex justify-between py-5 bg-white overflow-x-auto">
			{#each items as item, i}
				{#if item.count}
					<a class:text-white={item.slug == active} class:bg-tipplers-primary={item.slug == active} class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href={item.path} title={item.title} on:click={() => onClickFilterItem(item)}>
						{item.name}
					</a>
				{:else}
					<span class="px-4 py-2.5 text-tipplers-secondary">{item.name}</span>
				{/if}
			{/each}
		</nav>
	</div>
</div>
