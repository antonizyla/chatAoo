import type { Meta, StoryObj } from '@storybook/svelte';

import TextInput from './TextInput.svelte';

const meta = {
	title: 'Form/TextInput',
	component: TextInput,
	tags: ['autodocs'],
	argTypes: {
		placeholder: { control: 'text' },
		id: { control: 'text' },
		label: { control: 'text' },
		autocomplete: { control: 'boolean' },
		required: { control: 'boolean' },
		validate: { control: 'boolean' },
		name: { control: 'text' },
		regex: { control: 'text' },
		validationMessage: { control: 'text' }
	}
} satisfies Meta<TextInput>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
	args: {
		placeholder: 'Your name here..',
		id: 'name',
		label: 'Name',
		autocomplete: false,
		required: true,
		validate: true,
		name: 'name',
		validationMessage: 'Please enter a name containing only alphanumeric characters and spaces.'
	}
};

export const Secondary: Story = {
	args: {
		placeholder: 'Your email here..',
		id: 'email',
		label: 'Email',
		autocomplete: true,
		required: true,
		validate: true,
		name: 'email',
		regex: '^[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*@[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*$',
		validationMessage: 'Please enter a valid email address.'
	}
};
