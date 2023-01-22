<script lang="ts">
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import { setRandomInterval } from '$lib/utils/time';
	import { getRandomElements } from '$lib/utils/arrays';

	const imageBaseClassList = 'absolute w-full h-full cursor-pointer';
	const wrapperBaseClassList = 'select-none';

	export let imageClasses = '';
	export let wrapperClasses = '';
	export let limitRandom = 10;
	export let limitImages = 48;
	export let minDelay = 1000;
	export let maxDelay = 4000;
	export let images: string[];

	let selected: string[] = [];

	onMount(() => {
		const interval = setRandomInterval(
			() => {
				for (var i = 0; i < Math.floor(Math.random() * limitRandom); i++) {
					var randomIndex = Math.floor(Math.random() * selected.length);
					selected[randomIndex] = images[Math.floor(Math.random() * images.length)];
				}
			},
			minDelay,
			maxDelay
		);
	});

	$: selected = getRandomElements(images, limitImages);
	$: imageCompositeClasses = `${imageBaseClassList} ${imageClasses}`;
	$: wrapperCompositeClasses = `${wrapperBaseClassList} ${wrapperClasses}`;
</script>

{#if images.length}
	<div class={wrapperCompositeClasses} in:fade>
		{#each selected as image, i}
			<div class="relative overflow-hidden aspect-square" transition:fade={{ delay: i * 50 }}>
				{#key image}
					<img src={image} class={imageCompositeClasses} alt="Image #{i}" transition:fade={{ duration: 1000 }} />
				{/key}
			</div>
		{/each}
	</div>
{/if}
