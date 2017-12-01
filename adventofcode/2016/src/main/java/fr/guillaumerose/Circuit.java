package fr.guillaumerose;

import static java.util.stream.Collectors.toList;

import java.util.HashMap;
import java.util.Map;

import fr.guillaumerose.operators.Formula;

public class Circuit {
	private final Map<String, Formula> formulas;
	private final Map<String, Integer> cache = new HashMap<>();

	public Circuit(Map<String, Formula> formulas) {
		this.formulas = formulas;
	}

	public Integer valueFor(String variable) {
		return cache.computeIfAbsent(variable, var -> compute(var));
	}

	private Integer compute(String variable) {
		Formula formula = formulas.get(variable);
		return formula.compute(formula.variables().stream().map(var -> valueFor(var)).collect(toList()));
	}
}