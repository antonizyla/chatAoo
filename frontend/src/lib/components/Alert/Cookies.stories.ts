import type { Meta, StoryObj } from '@storybook/svelte';

import Cookies from './Cookies.svelte';

const meta: Meta = {
	title: 'Cookies',
	component: Cookies,
	tags: ['autodocs'],
	args: {
		storybook: { control: 'boolean' }
	}
} satisfies Meta<Cookies>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
	args: {
		storybook: true
	}
};
