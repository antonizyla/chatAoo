<script lang="ts">
	export let quote: string;
	export let author: string;

	export let centre: boolean = false;
</script>

<div
	class:text-center={centre}
	class="m-4 max-w-[40ch] rounded-md bg-accent p-6 shadow shadow-accent/60 drop-shadow-sm"
>
	<div class="flex flex-row justify-between">
		<div class="self-start pr-3 text-7xl">❝</div>
		<p class="my-3 !pr-0 text-lg">{quote}</p>
		<div class="self-end !pl-0 text-7xl">❞</div>
	</div>
	<p class="pr-4 pt-4 text-right">- {author}</p>
</div>
