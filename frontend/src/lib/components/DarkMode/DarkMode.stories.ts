import type { Meta, StoryObj } from '@storybook/svelte';

import DarkMode from './Darkmode.svelte';

const meta = {
	title: 'DarkMode',
	component: DarkMode,
	argTypes: {
		darkMode: {
			control: {
				type: 'boolean'
			}
		}
	}
} satisfies Meta<DarkMode>;

export default meta;
type Story = StoryObj<typeof DarkMode>;

export const Primary: Story = {
	args: {
		darkMode: false
	}
};
