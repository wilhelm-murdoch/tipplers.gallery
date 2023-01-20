<script lang="ts">
	import { lazy } from '$lib/utils/lazy';

	export let src: string;
	export let url: string;
	export let from: string = 'thebar.com';
	export let name: string;
	export let attribution: string;
</script>

<div class="relative shadow-sm rounded-xl overflow-hidden isolate max-w-5xl select-none bg-white h-64 background-fallback">
	<div class="overflow-hidden cursor-pointer bg-tipplers-primary/20">
		<img class="is-lazy hover:scale-110 ease-in-out duration-500 h-52 w-full object-cover object-center" use:lazy={'/' + src.replace('/original', '/thumbnail')} alt="" />
	</div>
	<span class="absolute bottom-4 left-1/2 transform -translate-x-1/2 -translate-y-1/2 p-1 ring-4 ring-white/50 rounded-full overflow-hidden bg-white z-10">
		<a href={url} target="_blank" rel="noreferrer" title={name}>
			<img class="h-6 w-6" src="/images/icons/sources/{from}.png" alt={attribution} />
		</a>
	</span>
	
	<div class="absolute bottom-0 w-full text-center text-xs p-3 pt-5 font-serif truncate">{name}</div>
</div>
