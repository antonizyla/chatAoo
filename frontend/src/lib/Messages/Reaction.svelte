<script lang="ts">
	import Tooltip from '$lib/Tooltip/Tooltip.svelte';
	import { type message, wsPayload, currentUser, users } from './messagesStores';

	export let reaction: { user_id: string; reaction: string };
	export let message: message;

	async function removeReaction() {
		const payload = {
			action_type: 'removeReaction',
			reaction_emoji: reaction.reaction,
			user_id: reaction.user_id,
			message_id: message.message_id
		};
		$wsPayload = JSON.stringify(payload);
	}

	let hoveredOn = false;
</script>

<div
	on:mouseover={() => {
		hoveredOn = true;
	}}
	on:mouseleave={() => {
		hoveredOn = false;
	}}
	class:hovered={hoveredOn && reaction.user_id === $currentUser}
>
	{#if reaction.user_id == $currentUser}
		{reaction.reaction}
	{:else}
		<Tooltip outer={reaction.reaction}>{$users[reaction.user_id]}</Tooltip>
	{/if}
	{#if hoveredOn && reaction.user_id === $currentUser}
		<button on:click={removeReaction}>X</button>
	{/if}
</div>

<style lang="postcss">
	.hovered {
		@apply bg-red-400 rounded-full px-2;
	}
</style>
