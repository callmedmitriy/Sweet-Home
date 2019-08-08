import React from 'react';
import { MainBlock } from '../components/MainBlock'

import './App.css';

class App extends React.Component {
	render() {
		return (
			<div className="wrapper">
				<MainBlock />
				{/*
					MainBlock
						WelcomeText
						Form
					DescriptionBlock
						DescriptionText
						CommonData
					ResultsBlock
						ResultData
					ResultTableBlock
						ResultFullTable
					+-Block
				*/}
			</div>
		)
	}
}

export default App;