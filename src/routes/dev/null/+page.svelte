<script lang="ts">
	export let data: any;
</script>

<div class="flex flex-col h-screen p-4 max-w-7xl mx-auto">
	<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4" id="result-cards">Design Elements</h3>
	<navigation class="prose">
		<ul>
			<li><a href="#result-cards">Result Cards</a></li>
			<li><a href="#collection-cards">Collection Cards</a></li>
		</ul>
	</navigation>

	<div class="container mx-auto max-w-7xl">
		<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4" id="result-cards">Result Cards</h3>
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
			{#each data.cocktails.slice(0, 12) as cocktail}
				<div class="relative shadow-sm rounded-b-2xl">
					<div class="flex bg-white shadow-md cursor-pointer max-h-64 overflow-hidden rounded-tr-2xl rounded-tl-2xl">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={cocktail.images[0].source_url} alt="" />
						<a href={cocktail.source_url} target="_blank" rel="noreferrer" title="">
							<span class="absolute top-64 left-1/2 transform -translate-x-1/2 -translate-y-1/2 p-1 ring-4 ring-white/50 rounded-full overflow-hidden bg-white">
								<img class="h-6 w-6" src="https://www.thebar.com/icons/apple-touch-icon.png" alt="" />
							</span>
						</a>
					</div>
					<div class="bg-white w-full text-center py-6 font-bold rounded-b-2xl font-serif">
						<h3 class="underline">{cocktail.name}</h3>
					</div>
				</div>
			{/each}
		</div>
	</div>

	<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4" id="collection-cards">Collection Cards</h3>
	<div class="container mx-auto max-w-7xl">
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 pb-4">
			<div class="relative shadow-md rounded-2xl bg-white overflow-hidden">
				<div class="grid grid-cols-2 cursor-pointer">
					<div class="overflow-hidden">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={data.cocktails[0].images[0].source_url} alt="" />
					</div>
					<div class="overflow-hidden">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={data.cocktails[3].images[0].source_url} alt="" />
					</div>
					<div class="overflow-hidden">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={data.cocktails[7].images[0].source_url} alt="" />
					</div>
					<div class="overflow-hidden">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={data.cocktails[11].images[0].source_url} alt="" />
					</div>
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

			<div class="relative shadow-md rounded-2xl bg-white overflow-hidden">
				<div class="grid grid-cols-2 cursor-pointer">
					<div class="overflow-hidden">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={data.cocktails[43].images[0].source_url} alt="" />
					</div>
					<div class="overflow-hidden">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={data.cocktails[33].images[0].source_url} alt="" />
					</div>
					<div class="overflow-hidden">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={data.cocktails[71].images[0].source_url} alt="" />
					</div>
					<div class="overflow-hidden">
						<img class="hover:scale-125 ease-in duration-500 w-full object-cover object-center" src={data.cocktails[138].images[0].source_url} alt="" />
					</div>
				</div>
				<div class="absolute flex -space-x-3 overflow-hidden left-1/2 transform -translate-x-1/2 -translate-y-1/2">
					<span class="hover:z-50 z-50 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
						<img src="https://i.pravatar.cc/64?10" alt="" />
					</span>
					<span class="hover:z-50 z-40 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
						<img src="https://i.pravatar.cc/64?9" alt="" />
					</span>
					<span class="hover:z-50 z-30 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
						<img src="https://i.pravatar.cc/64?8" alt="" />
					</span>
					<span class="hover:z-50 z-20 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
						<img src="https://i.pravatar.cc/64?7" alt="" />
					</span>
					<span class="hover:z-50 z-10 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden cursor-pointer">
						<img src="https://i.pravatar.cc/64?6" alt="" />
					</span>
					<span class="z-0 inline-block h-12 w-12 rounded-full border-4 border-white/50 overflow-hidden pt-2 bg-tipplers-secondary font-serif font-bold align-bottom text-center">11</span>
				</div>
				<div class="w-full text-center mt-2 py-6 font-bold rounded-b-2xl font-serif">
					<h3 class="underline text-xl"><a href="/dev/null" title="" class="hover:text-slate-700">Beginner</a></h3>
					<span class="font-sans font-normal text-xs text-slate-400 mr-3">43 recipes &middot; 16 sources</span>
				</div>
			</div>
		</div>
	</div>
</div>
