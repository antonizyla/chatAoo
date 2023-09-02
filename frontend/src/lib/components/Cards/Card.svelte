<script lang="ts">
	export let title: string;
	export let content: string;
</script>

<div
	class="flex min-w-[16rm] flex-col justify-stretch rounded-md bg-secondary-button/95 p-4 text-center shadow-sm shadow-secondary-button drop-shadow-sm"
>
	<div class="p-2 font-bold text-text">
		{title}
	</div>
	<div class="max-w-[40ch] p-2 text-text/80">
		{content}
	</div>
</div>
