import type { Meta, StoryObj } from '@storybook/svelte';

import DatePicker from './DatePicker.svelte';

const meta: Meta = {
	title: 'DatePicker',
	component: DatePicker,
	tags: ['autodocs'],
	args: {
		name: { control: 'text' },
		label: { control: 'text' },
		minVal: { control: 'date' },
		maxVal: { control: 'date' },
		value: { control: 'date' }
	}
} satisfies Meta<DatePicker>;

export default meta;
type Story = StoryObj<typeof DatePicker>;

export const Primary: Story = {
	args: {
		name: 'date',
		label: 'Date'
	}
};

export const Birthday: Story = {
	args: {
		name: 'birthday',
		label: 'Your Birthday',
		minVal: new Date(1970, 0, 1),
		maxVal: new Date()
	}
};
