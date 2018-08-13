import nltk
nltk.download('wordnet')
nltk.download('stopwords')
from textblob import Word
from nltk.corpus import stopwords


def getSemanticSimilarity(sentence1, sentence2):

	stops = stopwords.words('english')
	# sentence1 = "Then perhaps we could have avoided a catastrophe."
	# sentence2 = "We would perhaps then able prevent a disaster."
	sentence1Split = sentence1.split(" ")
	sentence2Split = sentence2.split(" ")

	for word1 in sentence1Split:
		if word1 in stops:
			sentence1 = sentence1.replace(' ' + word1 + ' ', ' ')

	for word2 in sentence2Split:
		if word2 in stops:
			sentence2 = sentence2.replace(' ' + word2 + ' ', ' ')


	set1 = []

	for word in sentence1.split(" "):
		word = Word(word)
		if word.synsets != []:
			set1 += word.synsets

			for synset in word.synsets:
				nyms = ['synonyms', 'hypernyms', 'hyponyms', 'meronyms', 'holonyms', 'part_meronyms', 'sisterm_terms', 'troponyms', 'inherited_hypernyms']
				for i in nyms:
				    try:
				    	if(getattr(synset, i)() != []):
				    		set1 += getattr(synset, i)()
				    except AttributeError as e: 
				       
				        pass


	set2 = []

	for word in sentence2.split(" "):
		word = Word(word)
		if word.synsets != []:
			set2 += word.synsets

			for synset in word.synsets:
				nyms = ['synonyms', 'hypernyms', 'hyponyms', 'meronyms', 'holonyms', 'part_meronyms', 'sisterm_terms', 'troponyms', 'inherited_hypernyms']
				for i in nyms:
				    try:
				    	if(getattr(synset, i)() != []):
				    		set2 += getattr(synset, i)()
				    except AttributeError as e: 
				   
				        pass


	return len(list(set(set1)&set(set2)))

# def similarityForDefs(definitions1, definitions2):
	
# 	numPairs = 0
# 	sumOfSimilarities1 = 0

# 	for definition1 in definitions1:
# 		for definition2 in definitions2:

# 			sumOfSimilarities1 += similarityForTwoDefs(definition1, definition2)
# 			numPairs += 1

# 	if(numPairs == 0):
# 		return 0
# 	else:
# 		averageOfSimilarities = sumOfSimilarities1/float(numPairs)
# 		return sumOfSimilarities1 


# def similarityForTwoDefs(definition1, definition2):

# 	sumOfSimiliarities = 0
# 	numPairs = 0

# 	for word1 in definition1.split():
# 		if word1 in stops:
# 			definition1 = definition1.replace(' ' + word1 + ' ', ' ')

# 	for word2 in definition2.split():
# 		if word2 in stops:
# 			definition2 = definition2.replace(' ' + word2 + ' ', ' ')

# 	for word1 in definition1.split():
# 		for word2 in definition2.split():

# 			wordSimilarity = 0
# 			word1 = Word(word1)
# 			word2 = Word(word2)

# 			if(word1.synsets == [] and word2.synsets != []):
# 				wordSimilarity = 0
			
# 			elif(word1.synsets != [] and word2.synsets == []):
# 				wordSimilarity = 0

# 			elif(word1.synsets == [] and word2.synsets == []):
# 				wordSimilarity = 1
# 			else:
# 				synset1 = word1.synsets[0]
# 				synset2 = word2.synsets[0]
# 				wordSimilarity = synset1.path_similarity(synset2)

# 			if wordSimilarity == None:
# 				sumOfSimiliarities += 0
# 			else:
# 				sumOfSimiliarities += wordSimilarity
			
# 			numPairs += 1
    
# 	averageOfSimilarities = sumOfSimiliarities/float(numPairs)

# 	return sumOfSimiliarities 
    


# sumOfSimiliarities = 0
# numPairs = 0

# for word1 in sentence1.split():
# 	for word2 in sentence2.split():
# 	    w1 = Word(word1)
# 	    w2 = Word(word2)

# 	    similarityForPair = similarityForDefs(w1.definitions, w2.definitions)
# 	    sumOfSimiliarities += similarityForPair
# 	    numPairs += 1

# print(sumOfSimiliarities) 










