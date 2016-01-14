package fr.guillaumerose;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class App {
	public static void main(String[] args) throws Exception {
		Circuit build = parse("/input.txt");
		System.out.println(build.valueFor("a")); // 956
	}

	private static Circuit parse(String resource) throws IOException {
		BufferedReader in = new BufferedReader(
				new InputStreamReader(CircuitParser.class.getResource(resource).openStream()));
		String line = null;
		CircuitParser parser = new CircuitParser();
		while ((line = in.readLine()) != null) {
			parser.parse(line);
		}
		return parser.build();
	}
}