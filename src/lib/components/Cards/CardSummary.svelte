<script lang="ts">
	import { lazy } from '$lib/utils/lazy';
	import { fade } from 'svelte/transition';

	export let cocktail: any;
</script>

<div in:fade class="bg-white rounded-xl shadow-sm overflow-hidden h-auto">
	<div class="grid grid-cols-1 md:grid-cols-4">
		<div class="relative background-fallback cursor-pointer overflow-hidden border-0 lg:border-r border-tipplers-primary/30 bg-tipplers-primary/20">
			<img class="is-lazy hover:scale-110 ease-in-out duration-500 aspect-square w-full h-auto md:w-full md:h-full object-cover object-center" use:lazy={'/' + cocktail.images[0].relative_path.replace('/original', '/medium')} alt="" />
			<span class="absolute bottom-2 left-2 rounded-md bg-black/50 px-2 py-1 text-white text-xs">
				<svg class="w-4 h-4 inline-block" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
					<path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2 2 6.477 2 12s4.477 10 10 10zM13.5 9.17a3 3 0 100 5.659" stroke-linecap="round" stroke-linejoin="round" />
				</svg>
				<a href={cocktail.meta.source_url} target="_blank" rel="noreferrer" title="" class="underline">{cocktail.meta.sourced_from}</a>
			</span>
		</div>

		<div class="md:col-span-3 p-4 pt-0">
			<h3 class="relative border-b border-tipplers-lightest my-4 pb-5">
				<span class="font-serif font-bold text-xl md:text-3xl">{cocktail.name}</span>
				<a class="absolute top-0 right-0 p-1 ring-4 ring-tipplers-secondary rounded-full overflow-hidden bg-white" href={cocktail.meta.source_url} target="_blank" rel="noreferrer" title="Check out the full recipe for {cocktail.name} at {cocktail.meta.sourced_from}.">
					<img class="h-5 w-5 md:h-8 md:w-8" src="/images/icons/sources/{cocktail.meta.sourced_from}.png" alt="" />
				</a>
			</h3>
			<div class="prose font-sans font-normal text-xs text-slate-400 mb-2">
				From <a href={cocktail.meta.source_url} target="_blank" rel="noreferrer" title="Check out the full recipe for {cocktail.name} at {cocktail.meta.sourced_from}.">{cocktail.meta.sourced_from}</a> &middot; {cocktail.ingredients.length} Ingredients &middot; {cocktail.instructions.length} Steps
				{#if cocktail.meta.standard_drinks}
					&middot {cocktail.meta.standard_drinks} {cocktail.meta.standard_drinks == 1 ? 'Unit' : 'Units'}
				{/if}
			</div>
			<div class="border-b border-tipplers-lightest mb-2.5">
				<p class="text-slate-600 prose md:prose-lg pb-2.5">{cocktail.description}</p>
			</div>

			<!--<div class="border-b border-tipplers-lightest pb-2.5 mb-2.5">
				<h3 class="relative prose font-serif font-bold text-lg cursor-pointer">
					Ingredients
					<svg class="absolute text-tipplers-secondary right-0 -top-0.5 h-8 w-8" fill="none" stroke="currentColor" stroke-width="3" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
						<path stroke-linecap="round" stroke-linejoin="round" d="M4.5 15.75l7.5-7.5 7.5 7.5" />
					</svg>
				</h3>
				<div class="border-t border-tipplers-lightest/50 mt-2">
					<div class="prose text-slate-500">
						<ul>
							{#each cocktail.ingredients as ingredient}
								<li>{ingredient.name}</li>
							{/each}
						</ul>
					</div>
				</div>
			</div>
			<div class="border-b border-tipplers-lightest pb-2.5 mb-2.5">
				<h3 class="relative prose font-serif font-bold text-lg cursor-pointer">
					Equipment
					<svg class="absolute text-tipplers-secondary right-0 -top-0.5 h-8 w-8" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
						<path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
					</svg>
				</h3>
				<ul>
					{#each cocktail.equipment as equipment}
						<li>{equipment.name}</li>
					{/each}
				</ul>
			</div>
			<div class="border-b border-tipplers-lightest pb-2.5 mb-5">
				<h3 class="relative prose font-serif font-bold text-lg cursor-pointer">
					Method
					<svg class="absolute text-tipplers-secondary right-0 -top-0.5 h-8 w-8" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
						<path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
					</svg>
				</h3>
				<ol>
					{#each cocktail.instructions as instruction}
						<li>{instruction.description}</li>
					{/each}
				</ol>
			</div>-->
			<div class="pb-2.5">
				{#each cocktail.tags as tag}
					{#if tag.name}
						<span class="text-xs pr-1.5">
							<span class="text-slate-400 pr-0.5">#</span>{tag.name}
						</span>
					{/if}
				{/each}
			</div>
		</div>
	</div>
</div>
