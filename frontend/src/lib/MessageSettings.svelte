<script lang="ts">
	import { messages, currentUser, type message } from '$lib/messagesStores';

	let open: boolean = false;

	export let message: message;

	async function deleteMessage() {
		const status = await fetch(`http://localhost:8081/messages/${message.message_id}`, {
			method: 'DELETE'
		}).then((res) => res.status);
		console.log(status);
		if (status === 200) {
			console.log('Message Deleted, updating ui');
			// find the message and set deleted equal to true
			$messages.forEach((msg) => {
				if (msg.message_id === message.message_id) {
					msg.deleted = true;
				}
				console.log(msg);
			});
		} else {
			console.log('Message not deleted due to internal server error');
		}
		$messages = $messages; // cause all messages to be rerendered
	}

    // logic to close the menu whenever the user scrolls
	let scrollPos: number;
	function toggle(_scrollpos: number) {
		open = false;
	}
	$: toggle(scrollPos);
</script>

<svelte:window bind:scrollY={scrollPos} />

<div
	class="absolute flex flex-col bg-red-400 origin-[top left] translate-y-6 p-3"
	class:hidden={!open}
>
	<ul class="list-none">
		<li><button>React</button></li>
		{#if $currentUser == message.user_id}
			{#if !message.deleted}
				<li><button>Edit</button></li>
				<li><button on:click={deleteMessage}>Delete </button></li>
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
