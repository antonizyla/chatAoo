import type { Meta, StoryObj } from '@storybook/svelte';

import Nav from './Nav.svelte';

const meta = {
	title: 'Nav',
	component: Nav,
	tags: ['autodocs'],
	argTypes: {
		title: { control: 'text', description: 'Company Name or Title' }
	}
} satisfies Meta<Nav>;

export default meta;
type Story = StoryObj<typeof Nav>;

export const primary: Story = {
	args: {
		mainLinks: [
			{ href: '#', label: 'Home' },
			{ href: '#', label: 'About' },
			{ href: '#', label: 'Contact' }
		],
		ctaLinks: [{ href: '#', label: 'Sign Up', primary: true }],
		title: 'My App'
	}
};
