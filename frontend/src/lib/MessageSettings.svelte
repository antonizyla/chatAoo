<script lang="ts">
	import { messages, currentUser, type message } from '$lib/messagesStores';

	let open: boolean = false;

	export let message: message;

	async function deleteMessage() {
		const status = await fetch(`http://localhost:8081/messages/${message_id}`, {
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
		}
	}
</script>

<button
	on:click={() => {
		open = !open;
	}}
>
	Options
</button>
{#if open}
	<button>React</button>
	{#if $currentUser == message.user_id}
		<button>Edit</button>
		<button>Delete Message</button>
	{/if}
{/if}
