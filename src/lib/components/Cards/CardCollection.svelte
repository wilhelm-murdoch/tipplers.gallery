<script lang="ts">
	import { lazy } from '$lib/utils/lazy';
	import { getImageUrl } from '$lib/utils/urls';

	export let cocktails: any[];
</script>

<div class="relative shadow-sm rounded-2xl bg-white overflow-hidden isolate">
	<div class="grid grid-cols-3 cursor-pointer">
		{#each cocktails as cocktail}
			<div class="overflow-hidden background-fallback bg-tipplers-primary/20">
				<img class="is-lazy hover:scale-110 ease-in-out duration-500 inset-0 aspect-square h-full w-full object-cover object-center" use:lazy={getImageUrl(cocktail.images[0].relative_path, "medium")} alt="" />
			</div>
		{/each}
	</div>
	<div class="absolute flex -space-x-3 overflow-hidden left-1/2 transform -translate-x-1/2 -translate-y-1/2">
		<span class="hover:z-50 z-50 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
			<img src="https://i.pravatar.cc/64?5" alt="" />
		</span>
		<span class="hover:z-50 z-40 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
			<img src="https://i.pravatar.cc/64?4" alt="" />
		</span>
		<span class="hover:z-50 z-30 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
			<img src="https://i.pravatar.cc/64?3" alt="" />
		</span>
		<span class="hover:z-50 z-20 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
			<img src="https://i.pravatar.cc/64?2" alt="" />
		</span>
		<span class="hover:z-50 z-10 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
			<img src="https://i.pravatar.cc/64?1" alt="" />
		</span>
		<span class="z-0 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden pt-2 bg-tipplers-secondary font-serif font-bold align-bottom text-center">27</span>
	</div>
	<div class="w-full text-center mt-2 py-6 font-bold rounded-b-2xl font-serif">
		<h3 class="underline text-xl"><a href="/dev/null" title="" class="hover:text-slate-700">Summer Cocktails</a></h3>
		<span class="font-sans font-normal text-xs text-slate-400 mr-3">343 recipes &middot; 32 sources</span>
	</div>
</div>
