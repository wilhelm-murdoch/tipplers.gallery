<script lang="ts">
	import { lazy } from '$lib/utils/lazy';

	export let cocktail: any;
</script>

<div class="relative shadow-sm rounded-2xl bg-white overflow-hidden isolate">
	<div class="flex cursor-pointer max-h-64 overflow-hidden background-fallback border-b border-tipplers-primary/30 bg-tipplers-primary/20">
		<img class="is-lazy hover:scale-110 ease-in-out duration-500  h-52 w-full  object-cover object-center" use:lazy={'/' + cocktail.images[0].relative_path.replace('/original', '/medium')} alt="" />
	</div>
	<a href={cocktail.source_url} target="_blank" rel="noreferrer" title="">
		<span class="absolute left-1/2 transform -translate-x-1/2 -translate-y-1/2 p-1 ring-4 ring-white/50 rounded-full overflow-hidden bg-white">
			<img class="h-6 w-6" src="/images/icons/sources/thebar.com.png" alt="" />
		</span>
	</a>
	<div class="w-full text-center pt-6 pb-3 font-bold font-serif">
		<h3 class="underline">{cocktail.name}</h3>
		<span class="font-sans font-normal text-xs text-slate-400 mr-3">{cocktail.ingredients.length} ingredients &middot; {cocktail.instructions.length} steps </span>
	</div>
</div>
