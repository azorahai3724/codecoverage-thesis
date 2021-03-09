import React from 'react';
import { render, screen } from '@testing-library/react';
import {AppForm} from './App';

test('renders learn react link', () => {
  render(<AppForm />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
