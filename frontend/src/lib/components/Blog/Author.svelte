<script lang="ts">
	import Image from '$lib/components/Images/Image.svelte';

	export let name: string;
	export let email: string;
	export let avatar: string;
	export let bio: string;
</script>

<div
	class="shadow-inner-md flex w-fit flex-row gap-5 rounded-lg bg-secondary-button p-4 shadow shadow-secondary-button/70 drop-shadow-lg"
>
	<div class="rounded-full p-3 shadow-text drop-shadow-xl">
		<Image pad={false} round={true} src={avatar} alt={`Image of ${name}`} />
	</div>
	<div class="flex max-w-sm flex-col justify-between gap-1 p-2 py-6">
		<h2 class="text-2xl underline underline-offset-[12px]">{name}</h2>
		<p>{bio}</p>
		<p class="px-6 text-right text-lg text-text/70">{email}</p>
	</div>
</div>
