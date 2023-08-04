<script lang="ts">
	import Tooltip from '$lib/Tooltip/Tooltip.svelte';
	import Reaction from './Reaction.svelte';
	import { type message, reactions, users } from './messagesStores';

	export let message: message;

	let hoveredOn = false;
	$: console.log(hoveredOn);
</script>

{#if $reactions[message.message_id]}
	<div class="flex flex-row justify-start gap-1">
		{#each $reactions[message.message_id] as reaction}
			<div class="p-0.5">
				<Reaction {message} {reaction} />
			</div>
		{/each}
	</div>
{/if}
