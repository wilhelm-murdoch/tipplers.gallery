<script lang="ts">
	import { lazy } from '$lib/utils/lazy';
	import { fade } from 'svelte/transition';

	const imageBaseClassList = 'is-lazy absolute w-full h-full cursor-pointer';
	const wrapperBaseClassList = 'select-none';

	export let imageClasses = '';
	export let wrapperClasses = '';
	export let images: string[] = [];

	$: imageCompositeClasses = `${imageBaseClassList} ${imageClasses}`;
	$: wrapperCompositeClasses = `${wrapperBaseClassList} ${wrapperClasses}`;
</script>

<div class={wrapperCompositeClasses}>
	{#each images as image, i}
		<div class="relative overflow-hidden aspect-square">
			<img use:lazy={image} class={imageCompositeClasses} alt="Image #{i}" />
		</div>
	{/each}
</div>

<style>
	img.is-lazy {
		opacity: 0;
		transition: all 1s ease;
	}
</style>
