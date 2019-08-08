import React from 'react';

import './Form.css';

export class Form extends React.Component {

	state = {

		roomCount: 'one-room',
		floatState: 'old',
		city: 'RND',
		mortagePercentage: '',
		accumulated: '',
		ownHouse: false,
		firstPay: false

	}

	handleRadioChange = (e) => {
		const { name } = e.currentTarget
		this.setState({ [name]: e.currentTarget.value })
	}

	handleSelectChange = (e) => {
		this.setState({ city: e.currentTarget.value })
	}

	handleTextChange = (e) => {
		const { id,value } = e.currentTarget
		this.setState({ [id]: value.replace(/[^0-9.]+/g,"") })
	}

	handleCheckBoxChange = (e) => {
		const { id } = e.currentTarget
		this.setState({ [id]: e.currentTarget.checked })
	}


	onBtnClickHandler = (e) => {
		console.log(this.state)
	}

	render() {
		return (
			<form className="form">
				<h2 className="form__header">Давайте посчитаем</h2>
				<hr className="form__hr"/>
				<div className="form__group"> 
					<div className="form__inline">
						<label className="form__label">Интересует: </label>
						<input 
							type="radio" 
							name="roomCount" 
							id="oneRoom" 
							value="one-room"
							onChange={this.handleRadioChange}
							checked={this.state.roomCount === "one-room"}/>
						<label className="form__label-radio" htmlFor="oneRoom">1-к</label>
						<input 
							type="radio" 
							name="roomCount" 
							id="twoRooms" 
							value="two-rooms"
							onChange={this.handleRadioChange}
							checked={this.state.roomCount === "two-rooms"}/>
						<label className="form__label-radio" htmlFor="twoRooms">2-к</label>
						<input 
							type="radio" 
							name="roomCount" 
							id="threeRooms" 
							value="three-rooms"
							onChange={this.handleRadioChange}
							checked={this.state.roomCount === "three-rooms"}/>
						<label className="form__label-radio" htmlFor="threeRooms">3-к</label>
						<label className="form__label">квартира</label>
					</div>
				</div>
				<div className="form__group"> 
					<div className="form__inline">
						<label className="form__label">Тип: </label>
						<input 
							type="radio" 
							name="floatState" 
							id="newFloat"
							value="new"
							onChange={this.handleRadioChange}
							checked={this.state.floatState === "new"}/>
						<label className="form__label-radio" htmlFor="newFloat">Новостройка</label>
						<input 
							type="radio" 
							name="floatState" 
							id="oldFloat"
							value="old"
							onChange={this.handleRadioChange}
							checked={this.state.floatState === "old"}/>
						<label className="form__label-radio" htmlFor="oldFloat">Вторичка</label>
					</div>
				</div>
				<div className="form__group">
					<label className="form__label">В городе</label>
					<select value={this.state.city} onChange={this.handleSelectChange} className="form__select">
						<option value="MSC">Москва</option>
						<option value="SPB">Санкт-Петербург</option>
						<option value="SOC">Сочи</option>
						<option value="RND">Ростов-на-Дону</option>
						<option value="ECB">Екатеринбург</option>
					</select>
				</div>
				<div className="form__group"> 
					<label className="form__label">Процент ипотеки</label>
					<input 
						type="text" 
						className="form__text" 
						placeholder="10 %"
						id="mortagePercentage"
						onChange={this.handleTextChange}
						value={this.state.mortagePercentage}/>
				</div>
				<div className="form__group"> 
					<label className="form__label">Накоплено (₽)</label>
					<input 
						type="text" 
						className="form__text" 
						placeholder="1 000 000 ₽"
						id="accumulated"
						onChange={this.handleTextChange}
						value={this.state.accumulated}/>
				</div>
				<div className="form__group">
					<div className="form__inline">
						<input 
							type="checkbox"
							id="ownHouse"
							onChange={this.handleCheckBoxChange}/>
						<label className="form__label-checkbox" htmlFor="ownHouse"> Наличие собственного жилья</label>
					</div>
				</div>
				<div className="form__group">
					<div className="form__inline">
						<input 
							type="checkbox"
							id="firstPay"
							onChange={this.handleCheckBoxChange} />
						<label className="form__label-checkbox" htmlFor="firstPay"> Пепвая покупка жилья</label>
					</div>
				</div>
				<hr className="form__hr"/>

				<span className="form__button" onClick={this.onBtnClickHandler}>Посчитать</span>
			</form>
		)
	}
}

