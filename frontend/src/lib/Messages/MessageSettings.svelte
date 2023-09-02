<script lang="ts">
	import {
		messages,
		currentUser,
		type message,
		wsPayload,
		reactions
	} from '$lib/Messages/messagesStores';

	let open: boolean = false;
	let reactionOpen: boolean = false;

	export let message: message;

	function delMessage() {
		console.log('Deleting message');
		const payload = {
			action_type: 'deleteMessage',
			message_id: message.message_id
		};
		$wsPayload = JSON.stringify(payload);
	}

	const emojis = ['👍', 'b', 'c'];

	// logic to close the menu whenever the user scrolls
	let scrollPos: number;
	function toggle(_scrollpos: number) {
		open = false;
	}
	$: toggle(scrollPos);

	function newReaction(emoji: string) {
		const payload = {
			action_type: 'addReaction',
			message_id: message.message_id,
			user_id: $currentUser,
			reaction_emoji: emoji
		};
		$wsPayload = JSON.stringify(payload);
	}
</script>

<svelte:window bind:scrollY={scrollPos} />

<div
	class=""
	on:mouseleave={() => {
		open = false;
	}}
	role="tooltip"
>
	<div
		class="absolute flex flex-col bg-red-400 origin-top-left translate-y-6 p-3"
		class:hidden={!open}
	>
		<ul class="list-none">
			<li
				on:mouseover={() => {
					reactionOpen = true;
				}}
				on:focus={() => {}}
			>
				<button>React</button>
				{#if reactionOpen}
					<div class="flex flex-row gap-2">
						{#each emojis as reaction}
							<button
								class="hover:scale-110"
								on:click={() => {
									console.log('Reacting with ' + reaction);
									open = false;
									newReaction(reaction);
								}}>{reaction}</button
							>
						{/each}
					</div>
				{/if}
			</li>

			{#if $currentUser == message.user_id}
				{#if !message.deleted}
					<li><button>Edit</button></li>
					<li><button on:click={delMessage}>Delete </button></li>
				{/if}
			{/if}
		</ul>
	</div>
	<button
		on:click={() => {
			open = !open;
		}}
	>
		Options
	</button>
</div>
