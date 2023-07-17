<script lang="ts">
	export let sender_name: string;
	export let content: string;
	export let time: string;
	export let alignRight: boolean = false;

	const timeStamp = new Date(time);
	let displayTime = timeStamp.toLocaleTimeString().slice(0, -3);
</script>

<div class="message" class:alignRight>
	<div class="title">
		{sender_name}
	</div>
	<div class="content">
		{content}
	</div>
	<div class="date">
		{displayTime}
	</div>
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

	.alignRight {
		text-align: right;
	}
</style>
