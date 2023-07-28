<script lang="ts">
	import { users, type message } from '$lib/messagesStores';

	export let message: message;

	const timeStamp = new Date(message.created_at);
	let displayTime = timeStamp.toLocaleTimeString().slice(0, -3);

	import MessageSettings from './MessageSettings.svelte';
</script>

<div class="message">
	<div class="title">
		{$users[message.user_id]}
	</div>
	<div class="content">
		{message.body}
	</div>
	<div class="date">
		{displayTime}
	</div>
	<MessageSettings {message} />
</div>

<style>
	.message {
		display: flex;
		flex-direction: column;
		background-color: #f0f0f0;
		border-radius: 10px;
		padding: 10px;
		margin: 10px;
	}

	.title {
		font-weight: bold;
	}

	.content {
		margin-top: 5px;
	}

	.date {
		font-size: 10px;
		margin-top: 5px;
	}
</style>
