<script lang="ts">
	import CardSummary from '$components/Cards/CardSummary.svelte';
	import { fade } from 'svelte/transition';
	import { onMount } from 'svelte';
	import type { Cocktail } from '$lib/utils/types';

	export let data: Cocktail[] = [];

	let selected: string[] = [];

	const cocktailSelected = (cid: string) => {
		return selected.find((v) => v == cid) ? true : false;
	};

	const removeSelected = (cid: string) => {
		var index = selected.indexOf(cid);
		if (index > -1) {
			selected.splice(index, 1);
			selected = selected;
		}
	};

	const selectCocktail = (cid: string) => {
		if (cocktailSelected(cid)) {
			removeSelected(cid);
			return;
		}

		selected.push(cid);

		selected = selected;
	};
</script>

<div class="grid grid-cols-5 h-screen">
	<div class="p-5 col-span-3 overflow-y-scroll">
		{#each data.cocktails as cocktail}
			<div class="mb-5">
				<CardSummary {cocktail} />
				<div class="text-right">
					<button
						class="mt-5 border p-2 bg-tipplers-primary rounded-md shadow-sm text-white"
						on:click={() => {
							selectCocktail(cocktail.cid);
						}}>Select</button
					>
				</div>
			</div>
		{/each}
	</div>
	<div class="bg-black text-white p-5 col-span-2">
		<h2 class="border-b border-slate-600 text-center mb-5">{selected.length} Selected</h2>
		{#each selected as cid}
			<div>{cid}</div>
		{/each}
	</div>
</div>
