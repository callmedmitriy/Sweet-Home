import React from 'react';
import { Form } from '../components/Form'

import './MainBlock.css';

export class MainBlock extends React.Component {
	render() {
		return (
			<div className="main-block">
				<div className="info-wrapper">
					<div className="info">
						<h1 className="info__header">Ипотека или аренда?</h1>
						<hr className="info__hr"/>
						<p className="info__description">Проснувшись однажды утром после беспокойного сна, Грегор Замза обнаружил, что он у себя в постели превратился в страшное насекомое.</p>
						<p className="info__description">Лежа на панцирнотвердой спине, он видел, стоило ему приподнять голову, свой коричневый, выпуклый, разделенный дугообразными чешуйками живот, на верхушке которого еле держалось готовое вот-вот окончательно сползти одеяло.</p>
						<div className="info__social">
							<div className="person">
								<img className="person__photo" src="img/photo.jpg" alt=""></img>
								<p className="person__name">Дмитрий Чеканов</p>
							</div>
						</div>
					</div>		
				</div>
				<Form />
			</div>
		)
	}
}

